# capy Discord bot
- Discord bot using Go and the Discord API for fun and servers with friends
- Also hosting it currently on AWS EC2 instance (until the free trial ends, then GCP Compute Engine probably)
- Jenkins CI/CD also set up

## Wanna try capy out?
[Use this link to add it to your server](https://discord.com/oauth2/authorize?client_id=1390895214096551967&permissions=1758097384131648&integration_type=0&scope=applications.commands+bot)

## Current Functionality:
- basic commands
    - !pet : nice way to thank capy
    - !react4role (role name) : capy creates a role with the inputted name and anyone who reacts to it gets assigned it; removing your reaction removes you from the role
    - !deleteRole (role name) : capy deletes the given role; only works if capy was the one who created the role
    - !fact : random fun fact
    - !dadJoke : random dad joke
    - !help : lists all the commands capy can do
    - !remindMe (days):(hours):(minutes) (optional message) : capy reminds the user after the given amount of time by replying to the message that calls this command with the @ enabled and the inputted message if any
    - !ra : picks a random attack operator from Rainbow Six Siege so you don't have to ask your squad who to play
    - !rd : picks a random defense operator from Rainbow Six Siege
    - !owt : picks a random tank from Overwatch
    - !ows : picks a random support from Overwatch
    - !owd : picks a random DPS from Overwatch

## Desired Future Functionality:
- slash commands? (at some point, would be nice for autocomplete and look more official)
- chess bot (low priority, maybe after completing other projects)
