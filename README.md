smtpsinkgo
==========

A little go server that will receive emails from programs but not deliver them, rather just log them. Great for testing code.



## What is SMTP Sink Go

A simple SMTP server that does not actually send out any email, rather it 
writes them out a log. This is meant to be a tool that developers can use to
safely test out the email functionality of their applications without having
to stand up a full SMTP server or otherwise risk having the emails go out to
actual recipients, who would likely be upset at having their mailboxes filled
with such emails.