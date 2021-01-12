<pre>
██████╗  █████╗ ██████╗    ███████╗██╗  ██╗███████╗██╗     ██╗     
██╔══██╗██╔══██╗██╔══██╗   ██╔════╝██║  ██║██╔════╝██║     ██║     
██████╔╝███████║██║  ██║   ███████╗███████║█████╗  ██║     ██║     
██╔══██╗██╔══██║██║  ██║   ╚════██║██╔══██║██╔══╝  ██║     ██║     
██████╔╝██║  ██║██████╔╝██╗███████║██║  ██║███████╗███████╗███████╗
╚═════╝ ╚═╝  ╚═╝╚═════╝ ╚═╝╚══════╝╚═╝  ╚═╝╚══════╝╚══════╝╚══════╝
</pre>
### A PHP Web Shell made with ♥

---

**Description:**

This PHP Web Shell was made to have my own PHP Shell for TryHackMe Challenges, and for learn basic PHP syntax.\
For now is usable and in a stable version, but still in development.

GoShell™ is a Golang reverse-shell for have a stable and initial shell to a compromised host.\
(or almost the initial idea was that...)

Under /src folder there is:
- goshell.go [ the shell golang source code ]
- goshell [ the compiled binary of go code ]
- build.sh [ a minimal sh script to build the shell, then pack with UPX Packer ]

Then after that i encode the binary in base64 for better integration in the .php file.

---

**Usage:**

Simply use any exploit technique you found to upload the main file who was *b4dsh.php* into the webserver.\
Then point your browser to the uploaded .php shell, and you are ready to go!

---

**Disclaimer:**

*Compromising other people's computer systems without permission is a crime punishable by law.*\
*This software must be used only and exclusively for educational purposes.*\
*I do not hold myself responsible in any way or circumstance for the improper or illegal use of this software.*

**Fabio Corona - 2021**
