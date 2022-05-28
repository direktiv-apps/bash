Feature: greeting end-point

Background:
* url demoBaseUrl

Scenario: base

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    { 
        "commands": [
            {
                "command": "ls -la"
            }
        ]
    }
    """
    When method post
    Then status 200

Scenario: script

    Given path '/'
    Given header Direktiv-ActionID = 'development'
    Given header Direktiv-Tempdir = '/tmp'
    And request 
    """
    {   
        "files": [
            {
                "name": "run.sh",
                "mode": "0755",
                "data": "#!/bin/bash\necho HELLO"
            }
        ],
        "commands": [
            {
                "command": "./run.sh"
            }
        ]
    }
    """
    When method post
    Then status 200