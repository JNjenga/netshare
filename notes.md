## Client cli commands

1. ls

Returns a list of indexed files

`ls <path>`

2. cd 

Change current view directory

`cd <path>`

3. cp srcpath destpath

Downlaod file specified in srcpath to destpath

`cp <srcpath> <destpath>`

4. exit

Close application

## Protocol

1. ls

Request:

`ls <path>`

Respnose

```
<list_len>\n
<path_char_len><file_name>\n
<path_char_len><file_name>\n
```

2. cd

Request:

`cd <path>`

Response:

`<path>`

3. cp

Request:

`cp <srcpath> <destpath>`

Response:

```
<file_count>\n

<file_name>\n
<file_len>\n
<bytes>

```
