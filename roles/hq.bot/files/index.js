"use strict";

const express = require("express");
const bodyParser = require("body-parser");
const { spawn } = require("child_process");
const { WebClient } = require("@slack/web-api");
const { createEventAdapter } = require("@slack/events-api");
const { createMessageAdapter } = require("@slack/interactive-messages");

let targets = ["isu1.sysad.net", "isu2.sysad.net", "isu3.sysad.net"];

const { slackToken, slackSigningSecret } = require("./credential.json");

const web = new WebClient(slackToken);
const slackEvents = createEventAdapter(slackSigningSecret);
const slackInteractions = createMessageAdapter(slackSigningSecret);

const githubEventHandler = async (req, res) => {
	const {ref, compare, pusher: {name}} = req.body;

	if(ref !== "refs/heads/master"){
		return res.sendStatus(200);
	}

	await web.chat.postMessage({
		channel: "CMNEKS1HR",
		text: `${name} pushed new commits to master.`,
		blocks: [
			{
				type: "section",
				text: {
					type: "mrkdwn",
					text: `${name} pushed new commits to master. <${compare}|View changes>\nWould you like to deploy them?`,
				},
			},
			{
				type: "actions",
				elements: [{
					action_id: "deploy",
					type: "button",
					style: "primary",
					text: {
						type: "plain_text",
						text: "Deploy",
					},
				}, {
					action_id: "skip",
					type: "button",
					text: {
						type: "plain_text",
						text: "Skip",
					},
				}],
			},
		],
	});

	return res.sendStatus(200);
};

slackEvents.on("app_mention", async event => {
	let text = "say \"<@kiritan> deploy [commit-id]\"";

	const [, op, arg] = event.text.split(" ");
	if(op === "ping"){
		text = "pong";
	}else if(op == "deploy"){
		deploy(arg);
		text = "Deployment process has been started. For details, see <#CM7SYH011>";
	}else if(op === "target"){
		targets = arg.split(",");
		text = `Targets changed. \`${JSON.stringify(targets)}\``;
	}

	await web.chat.postMessage({text, channel: event.channel});
});

slackInteractions.action({actionId: "deploy"}, async (payload, respond) => {
	deploy("origin/master");
	await respond({
		text: "Deployment process has been started.",
		blocks: [payload.message.blocks[0], {
			type: "context",
			elements: [{
				type: "mrkdwn",
				text: "🚀 Started. For details, see <#CM7SYH011>",
			}],
		}],
	});
});
slackInteractions.action({actionId: "skip"}, async (payload, respond) => {
	await respond({
		text: "Deployment process has been skipped.",
		blocks: [payload.message.blocks[0], {
			type: "context",
			elements: [{
				type: "mrkdwn",
				text: "🆗 Skipped. When you want to deploy some commit, say \"<@kiritan> deploy [commit-id]\"",
			}],
		}],
	});
});

const deploy = ref => Promise.all(targets.map(target => new Promise((resolve, reject) => {
	// The code below is vulnerable to OS command injection. Use carefully.
	const child = spawn("ssh", ["-f", target, `bash -c ". /etc/profile; make REF=${ref} 2>&1 | notify_slack"`]);
	child.on("exit", resolve);
	child.on("error", reject);
})));

const app = express();
app.use("/gh", bodyParser.json());
app.post("/gh", githubEventHandler);
app.post("/event", slackEvents.requestListener());
app.post("/interactive", slackInteractions.requestListener());
console.log("listening on 5000");
app.listen(5000);
