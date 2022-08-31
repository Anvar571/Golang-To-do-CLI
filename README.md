# TO-DO

1. Todo list 
2. Get todo by id, name, key
3. Create todo
4. Delete todo by id or name
5. Edit todo by id
6. Read todo by id, name
7. Change status

### to-do structure
- id - int
- title - string
- description - string
- status - string
- created_at time

### Flags
- list - return todo list
- get {id} return todo by id
- create {title} {discription}
- edit {id} {title} {discription}
- status {id} {status} - return ok

### file format
{id};{title};{discription};{status};{created_at}

### statuses
0 - todo
1 - in-progres
2 - done
3 - stopped 
