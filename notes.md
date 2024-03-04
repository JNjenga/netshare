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

### Protocol Format

nsv1 protocol, is a simple stateless http like protocol

## Request

```
<request_byte_size><command>\n
<arg 1>\n
<arg 2>\n
<arg .>\n
<arg .>\n
<arg n>\n
```

## Response

```
<response_byte_size><status_code>
<response_data>
```

`<response_byte_size>`: unsigned 32 bit integer

`<status_code>`: unsigned char(1 byte). Http like status codes

    - 0b00001--- - Success
    - 0b0001---- - Request/Client error
    - 0b001----- - Server error
    - 0b1------- - Custom

### Commands

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
