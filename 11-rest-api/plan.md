# Rest Api

Required routes:

|Http Verb|Route|Description|Authentication Required| Only Creator|Notes|
|---------|------|----------|--|------|-----|
| GET |     `/events`|          Get a list of available events||
| GET |     `/events/<id>`|     Get a event||
| POST |    `/events/`|         Create a new bookable event|YES||
| PUT |     `/events/<id>`|     Update an event|YES|YES|Requires JWT|
| DELETE |  `/events/<id>`|     Delete an event|Yes|YES|Requires JWT|
| POST |    `/signup`|          Create new user||
| POST |    `/login`|           Authenticate user|||Returns JWT|
| POST |    `/events/<id>/register`|   Register user for event|Yes||Requires JWT|
| DELETE |  `/events/<id>/register`|   Cancel registration|Yes||Requires JWT|
