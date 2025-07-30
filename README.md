# capy Discord bot
- Discord bot using Go and the Discord API just for fun. 
- Also hosting it currently on AWS EC2 instance (until the free trial ends, then GCP Compute Engine probably).
- Jenkins CI/CD also set up. 

## Current Functionality:
- basic commands
    - !pet : nice way to thank capy
    - !react4role <role name> : capy creates a role with the inputted name and anyone who reacts to it gets assigned it; removing your reaction removes you from the role
    - !deleteRole <role name> : capy deletes the given role; only works if capy was the one who created the role
        - TODO: currently the roles capy made are stored in memory so if capy goes down (like when changes are pushed), it'll forget what roles it made before and will refuse to delete those. Currently working on storing the roles in a DB to fix this.
        - TODO: add thread safety
    - !fact: random fun fact
        - TODO: add more APIs or have ChatGPT make a huge list
        - TODO: some sort of fact check pipeline? would be hard since LLM APIs aren't generous with free tiers if at all
    - !help: lists all the commands capy can do
    - !remindMe <days>:<hours>:<minutes> : capy reminds the user after the given amount of time by replying to the message that calls this command with the @ enabled
        - TODO: 
        

## Desired Future Functionality:
- chess bot (low priority, maybe after completing other projects)
