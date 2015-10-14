# craig-o-mation
CalHacks 2015

###Inspiration
As technology advances, our society is harnessing technology to acquire goods in more convenient ways than ever. With the rise of services such as Postmates, Caviar, Diner Dash, etc., it is now possible to have most basic necessities delivered in a timely fashion. However, one industry that appeals to us is the second hand goods industry. Second hand goods are often a good way to save money while getting similar quality products. One of our favourite sites is Craigslist, and we knew we could improve upon the buying experience.

###What it does
After a user has found a desirable product on Craigslist, he/she copies and pastes the URL into craig-o-mation. craig-o-mation will scrape the Craigslist post for information such as cost, and also requests a bit more necessary information from the purchaser. After the purchaser fills out this short form, the poster gets an email describing what the purchaser has offered. Should they choose to accept, after entering some information, the sale is completed. A Postmates courier is instantly dispatched to the purchaser to collect the product and deliver it to the buyer. A payment is made from the purchaser to the buyer through Capital One's APIs. Within minutes, the buyer gets his second hand goods conveniently for an affordable price.

###How we built it
- The forms are all written in HTML, CSS, and Javascript
- Forms send data to our backend written in Go
- Go backend proxies through a Python server running at home in order to bypass Craigslist's imposed restrictions on certain IPs (they are quite strict)
- Go backend talks to Postmates and Capital One APIs to facilitate the transaction and delivery of the item
- Email and datastore are provided with App Engine, where the app is hosted
- Also use Google Maps API for frontend forms to ease entering of addresses

###Challenges we ran into
Craigslist is extremely strict about bots and automation. They have IP blocked many large cloud computing companies and have imposed other restrictions to make scraping and automation extremely challenging. We were able to get a home server running that we could proxy requests through in order to scrape data from Craigslist.

###Accomplishments that we're proud of
- The UI is simple, helpful, and intuitive
- The emails have styling beyond basic text emails
- The Craigslist proxy was challenging to implement and works quite well
- The backend is very efficient due to design decisions
- We were successfully able to integrate with Postmates and Capital One APIs

###What we learned
- Get it done mentality. There is always a way to make novel innovations, it just takes perseverance
- Some problems stumped us for many hours but in the end we did not need to compromise on functionality

###What's next for craig-o-mation
- Write browser extensions to store user data and be accessible while browsing Craigslist to decrease friction with users
- Faster and more reliable solutions for scraping Craigslist
- Rating system for sellers based on reviews

##Contributors
- [Josh Chorlton](http://joshchorlton.com)
- [Matthew Du](http://matthewdu.com)
- [Julia Cao](http://github.com/JuliaCao)

