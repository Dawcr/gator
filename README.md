Gator is a guided project of Boot.Dev  

# How to install:  

Requirements before installing:  

- Postgres  
- Go  
- Git

To install run:  
```
git clone https://github.com/Dawcr/gator.git && cd gator && go build  
```  

before running create a file called ~/.gatorconfig.json in your home directory containing:  
``` 
{
"db_url": "postgres://example"
}
```  
substitute postgres://example with your postgres connection string  


# Available commands:  

register \<name>  
    creates a new user, also logs it in  

login \<name>  
    login using an existing user  

users  
    print the list of all registered users  

addfeed \<name> \<url>  
    create a new feed entry, and follows it using the currently logged in user  

feeds  
    lists all feeds that have been added  

follow \<url>  
    follow a feed that has been previously added  

unfollow \<url>  
    unfollow a feed  

following  
    lists all feeds followed by logged in user  

agg \<time_between_requests>  
    aggregates the rss feed every specified interval, interval format: 1s, 1m, 1h, minimum is 1m  

browse \[limit]  
    browse aggregated posts by most recent, can specify a limit to display, default is 2  

reset  
    debug command, resets the database.  

# Example usage:  
```
./gator register daw
./gator addfeed "Boot.dev Blog" https://blog.boot.dev/index.xml
./gator agg 1h
```
switch to another terminal while this aggregates in background
```
./gator browse 10
```