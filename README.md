# capy Discord bot
- Discord bot using Go and the Discord API just for fun. 
- Also hosting it currently on AWS EC2 instance (until the free trial ends, then GCP Compute Engine probably).
- Jenkins CI/CD also set up. 

## Current Functionality:
- runs
- basic commands
    - !pet : nice way to thank capy
    - !react4role <role name> : capy creates a role with the inputted name and anyone who reacts to it gets assigned it; removing your reaction removes you from the role
        - the role is just made with permission to view channels, currently always is assigned the same color, might change to random or an inputted one in the future
    - !deleteRole <role name> : capy deletes the given role; only works if capy was the one who created the role
        - TODO: currently the roles capy made are stored in memory so if capy goes down (like when changes are pushed), it'll forget what roles it made before and will refuse to delete those. Currently working on storing the roles in a DB to fix this.
        - TODO: currently uses global variables which probably isn't best. Should try using pipes and adding some thread safety maybe.

## Desired Future Functionality:
- random fun facts
- "remind me" after a specified amount of time
- !help to list commands
- chess bot (maybe after completing other projects)
