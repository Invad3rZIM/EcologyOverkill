# EcologyOverkill

# Part I - What is this & why did you write it? 


Great questions! This is a tool I designed to get out of my homework. I graduate university in about 3.5 weeks, but at the moment I'm a part-time student who took Ecology as a science elective. In about 4 hours, my term-project is due. This project required collecting 15 different leaves, identifying them,
then researching their species classification ( Kingdom > Phylum > Class > Order > Family > Species ) and documenting it in a table.

This seemed like a lot of repetitive, boring work, and so naturally I ignored it until the very last day. Instead, I came across a job form for a tech-company that asked for "a sample of your code". Unfortunately for me, most of the code I write is either professionally done (and so I can't freely share it) or it's way too dated. So I decided I would find a mini-project to work on, & write some code using that company's tech stack.


See files/output.txt for raw final output.

Alternatively, check out finaltable.png (or scroll down) to see how how it turned out

# Part II - The Wiki Web Crawler
Then I got an idea - this ecology project is a lot of wikipedia lookups, and it's mindnumbing. Maybe I could partially automate my term-project. So I built a simple webcrawler this morning, one that reads from a list of websites (see files/crawllist.txt), sends a GET requests, then parses the body to find all the relevant data. I saved this data to files/filtereddata.txt.

Unfortunately (and perhaps expectedly), I can't make a flawless webcrawler in the course of an hour or 2. 

There are a couple problems

1. Not all the information I need can be found on wikipedia (about 90% of the data got scraped though). In particular, I needed to use a secondary site to get Phylum & Class information about several different species. No worries, though

2. I still need to check over my data, ensuring accuracy. This is a term-project after all, so I need to be thorough. :)

With that in mind, it made sense to write to files/filtereddata.txt

THat document contained my auto-generated code. I left this intact so you can see the end result of my web-crawler. 

files/modifieddata.txt includes changes I made to correct for data that wasn't on the wikipedia pages. The biggest problem was simply I chose wikipedia pages that didn't have all the requisite info. My bad!

Here's an example of filtereddata.txt vs modifieddata.txt

# Auto-generated (filtered)
```
Kingdom Plantae
Clade Tracheophytes
Clade Angiosperms
Clade Eudicots
Clade Rosids
Order Rosales
Family Rosaceae
Genus Prunus
Subgenus Prunus
```
# Modified (user-validated)
```
Kingdom Plantae
Clade Tracheophytes
Clade Angiosperms
Clade Eudicots
Clade Rosids
Order Rosales
Family Rosaceae
Phylum Magnoliophyta
Class Magnoliopsida
Genus Prunus
Subgenus Prunus
Species Prunus cerasifera, cherry plum
```
As far as I'm concerned, this web-crawler was a success!

# Part III - File Parser / Table Output

So at this point, I have all the data I need organized and scannable. Great!

Next is the fun(?) part! My professor wrote the following in the instructions...

```
3.   Results.  This will consist of a table that outlines the complete taxonomy and common name for each of your specimens.  Use the same phylogenetic sequence that was published in the main reference text you consult.  Each taxonomic group will be listed only once, and the species will be numbered consecutively from 1-15.  Use these same numbers, based on phylogenetic sequence, on the specimen labels (i.e., assemble the table first, then add those numbers to the labels).  

Please use the following format EXACTLY, including sequence and indentations: 

Kingdom:          Animalia
  Phylum:         Arthropoda
    Class:        Insecta
      Order:      Odonata
        Family:   Coenagrionidae
                    1.  Enallagma civile, civil bluet
                    2.  Argia americana, American forktail
      Family:     Aeshnidae
                    3.  Anax junius, green darner
    Order:        Lepidoptera
      Family:     Saturniidae
                    4.  Hyalophora cecropia, cecropia moth
      Family:     Danaidae
                    5.  Danaus plexippus, monarch butterfly
    Order:        Hymenoptera
      Family:     Apidae
                    6.  Megabombus pennsylvanicus, bumblebee
                    7.  Apis mellifera, honeybee
```

Ok, he wants the table *exact*. That's fair, this is science after all...

I can do that, but the first step is to organize my data.

Something to note is the number system. My ecology professor doesn't know it (or care), but the output he wants us to match exactly follows the form of an inorder-traversal DFS. Wonderful recursion, I know how to do that! It's worth mentioning he's got some interesting formatting stuff going on for the headers (Kingdom, Phylum, etc...) Not too big a problem.

Once I saw it was a DFS, the only real challenge was to get my modifiedData into a form that could be graphically traversed. For this, I went with an adjacency list.

It goes a little like this...

Every Phylum P is a child of a kingdom K
Every Class C is a child of a Phylum P
and so on...

Given that, all I had to do was throw everything into a map[string]map[string]bool to represent this.
The reason the value is a map[string]bool instead of a []string{} slice is that I didn't want repeats, so
I wanted to use the map as a set (I believe go only supports sets in this way). See engine.deriveRelationshipsSet() to full implementation.

At this point, I had all the pieces built, and all that was left was superficial formatting.

I did run into a single bug however with this approach, and that was that there was a single instance where 2 distinct Phylums had the same Class child. This meant that the adjacency list would render a species twice in that one instance. Because the scope of this project was simply to replace my term-project, I decided that instead of refactoring the entire thing to be bugless, it was fairly appropriate to hard-code a bug-patch in 3 lines of code (engine/renderTree.go, 96-98). All this code does is ignore that lone edgecase.

If I had to do this over knowing may repeat across Phylum's, I'd make my key Parent => Parent:Child and my value be Child => Child:Grandchild. This solution feels like structuring composite keys in a relational-database. Because the keys and values, just got more specific, they would likely all be unique, removing this bug.

Here's how it turned out...

![Alt text](finaltable.png?raw=true "Output")

# Part IV - How long did this take, What did you learn, & was it worth it?

This took way longer than actually doing the term-project as it was meant to be done. Probably about 5x longer actually. But that's alright, because I also built myself a code-sample resume-piece in the process. And most importantly, it was a great challenge and way better than doing it manually.

I don't think I learned anything specifically new, but I did get more experience with Go engineering, and that's always desirable. 

In conclusion, yes, it was absolutely worth it.

# Part V - Thank You For Reading!

Hey, thank you so much for getting to the end of this. 

If you liked this mini-project and are hiring in NYC or California, I'm totally up for an interview!


My name's Kirk Zimmer (kzimmer655@gmail.com)

My favorite languages are Go, Java, Python, Javascript (React & React-Native)

Shoot me an email!
