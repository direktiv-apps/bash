
Feature: Basic

# The secrects can be used in the payload with the following syntax #(mysecretname)
Background:

Scenario: simple

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{
		"commands": [
		{
			"command": "bash -c 'echo $Key'",
			"silent": true,
			"print": false,
			"envs": [
				{	
					"name": "Key",
					"value": "Value"
				}
			]
		}
		]
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	{
	"bash": [
	{
		"result": "Value",
		"success": true
	}
	]
	}
	"""
	
Scenario: file

	Given url karate.properties['testURL']

	And path '/'
	And header Direktiv-ActionID = 'development'
	And header Direktiv-TempDir = '/tmp'
	And request
	"""
	{	
		"files": [
			{
				"name": "run.sh",
				"data": "#!/bin/bash\necho TEST",
				"mode": "0755"
			}
		],
		"commands": [
		{
			"command": "./run.sh",
			"silent": true,
			"print": false,
		}
		]
	}
	"""
	When method POST
	Then status 200
	And match $ ==
	"""
	{
	"bash": [
	{
		"result": "TEST",
		"success": true
	}
	]
	}
	"""