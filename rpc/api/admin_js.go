package api

const Admin_JS = `
web3._extend({
	property: 'admin',
	methods:
	[
		new web3._extend.Method({
			name: 'addPeer',
			call: 'admin_addPeer',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'exportChain',
			call: 'admin_exportChain',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'importChain',
			call: 'admin_importChain',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'sleepBlocks',
			call: 'admin_sleepBlocks',
			params: 2,
			inputFormatter: [null, null]
		}),
		new web3._extend.Method({
			name: 'verbosity',
			call: 'admin_verbosity',
			params: 1,
			inputFormatter: [web3._extend.utils.fromDecimal]
		}),
		new web3._extend.Method({
			name: 'setSolc',
			call: 'admin_setSolc',
			params: 1,
			inputFormatter: [null]
		}),
		new web3._extend.Method({
			name: 'startRPC',
			call: 'admin_startRPC',
			params: 4,
			inputFormatter: [null, null, null, null]
		}),
		new web3._extend.Method({
			name: 'stopRPC',
			call: 'admin_stopRPC',
			params: 0,
			inputFormatter: []
		}),
		new web3._extend.Method({
			name: 'setGlobalRegistrar',
			call: 'admin_setGlobalRegistrar',
			params: 2,
			inputFormatter: [null,null]
		}),
		new web3._extend.Method({
			name: 'setHashReg',
			call: 'admin_setHashReg',
			params: 2,
			inputFormatter: [null,null]
		}),
		new web3._extend.Method({
			name: 'setUrlHint',
			call: 'admin_setUrlHint',
			params: 2,
			inputFormatter: [null,null]
		}),
		new web3._extend.Method({
			name: 'saveInfo',
			call: 'admin_saveInfo',
			params: 2,
			inputFormatter: [null,null]
		}),
		new web3._extend.Method({
			name: 'register',
			call: 'admin_register',
			params: 3,
			inputFormatter: [null,null,null]
		}),
		new web3._extend.Method({
			name: 'registerUrl',
			call: 'admin_registerUrl',
			params: 3,
			inputFormatter: [null,null,null]
		}),
		new web3._extend.Method({
			name: 'startNatSpec',
			call: 'admin_startNatSpec',
			params: 0,
			inputFormatter: []
		}),
		new web3._extend.Method({
			name: 'stopNatSpec',
			call: 'admin_stopNatSpec',
			params: 0,
			inputFormatter: []
		}),
		new web3._extend.Method({
			name: 'getContractInfo',
			call: 'admin_getContractInfo',
			params: 1,
			inputFormatter: [null],
		}),
		new web3._extend.Method({
			name: 'httpGet',
			call: 'admin_httpGet',
			params: 2,
			inputFormatter: [null, null]
		})
	],
	properties:
	[
		new web3._extend.Property({
			name: 'nodeInfo',
			getter: 'admin_nodeInfo'
		}),
		new web3._extend.Property({
			name: 'peers',
			getter: 'admin_peers'
		}),
		new web3._extend.Property({
			name: 'datadir',
			getter: 'admin_datadir'
		}),
		new web3._extend.Property({
			name: 'chainSyncStatus',
			getter: 'admin_chainSyncStatus'
		})
	]
});
`