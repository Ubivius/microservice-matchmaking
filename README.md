# microservice-matchmaking
Matchmaking microservice for our online game framework.

## Matchmaking endpoints

`GET` `/queue/{user_id}` Returns true if the user is in queue. `user_id=[string]`

`GET` `/health/live` Returns a Status OK when live.

`GET` `/health/ready` Returns a Status OK when ready or an error when dependencies are not available.

`POST` `/queue` Add new user to the queue. </br>
__Data Params__
```json
{
  "user_id": "string, required",
  "user_ip": "string, required",
}
```

`DELETE` `/queue/{user_id}` Delete user from queue.  `user_id=[string]`
