RELEASE NOTES for queuemetre - A command for getting the number of email queues
===================================================================================================
- "https://github.com/azumakuniyuki/queuemetre"

v1.0.6
---------------------------------------------------------------------------------------------------
- release: "Mon, 17 Jun 2024 18:07:05 +0900 (JST)"
- version: "1.0.6"
- changes:
  - Check the user owns the daemon process (OpenSMTPD)

v1.0.5
---------------------------------------------------------------------------------------------------
- release: "Wed, 12 Jun 2024 22:00:00 +0900 (JST)"
- version: "1.0.5"
- changes:
  - Fix the first character of the queue file in DMA:Dragonfly Mail Agent

v1.0.4
---------------------------------------------------------------------------------------------------
- release: "Wed, 12 Jun 2024 18:32:08 +0900 (JST)"
- version: "1.0.4"
- changes:
  - Support DMA: Dragonfly Mail Agent
  - DRY: Tiny code improvement
  - Add and call currenttimes(), systemvalues() subroutines
  - Print the list of supported MTAs at `--help` screen
  - Fix bug in v1.0.2, 1.0.3

v1.0.1
---------------------------------------------------------------------------------------------------
- release: "Thu, 25 Apr 2024 20:22:22 +0900 (JST)"
- version: "1.0.1"
- changes:
  - Code improvement on the `ps` command and its option
  - Tiny code improvements

v1.0.0
---------------------------------------------------------------------------------------------------
- release: "Tue,  9 Apr 2024 22:22:22 +0900 (JST)
- version: "1.0.0"
- changes:
  - Organized and released a disposable command to display the number of email queues written about
    10 years ago.
  - As of present, Sendmail, Postfix, and OpenSMTPD are supported.

