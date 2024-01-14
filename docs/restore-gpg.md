Restore
Import the private key.

➜ gpg --import-options restore --import private.gpg
gpg: key C8DE632E9A8A0BDD: public key "Ponder Stibbons <ponder.stibbons@unseen.edu>" imported
gpg: key C8DE632E9A8A0BDD: secret key imported
gpg: Total number processed: 1
gpg:               imported: 1
gpg:       secret keys read: 1
gpg:   secret keys imported: 1
This invocation imports the key from the file private.gpg in the current directory. The import option restore imports all necessary data for GnuPG to fully restore the key. The import option keep-ownertrust keeps the owner trust of the key instead of clearing it’s trust value. This saves having to manually set the trust value for the key later.

Enter the private key’s passphrase in the [Import Passphrase Prompt] to import the key.

Passphrase Prompt to Import GPG Key
Import Passphrase Prompt
Now, edit the freshly imported key.

➜ gpg --edit-key ponder.stibbons@unseen.edu
gpg (GnuPG) 2.2.20; Copyright (C) 2020 Free Software Foundation, Inc.
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.

Secret key is available.

sec  rsa4096/C8DE632E9A8A0BDD
     created: 2020-11-13  expires: never       usage: SC
     trust: unknown       validity: unknown
ssb  rsa4096/DBCD8B98F2F9188C
     created: 2020-11-13  expires: never       usage: E
[ unknown] (1). Ponder Stibbons <ponder.stibbons@unseen.edu>
Enter trust to modify the trust value of the key.

gpg> trust
sec  rsa4096/C8DE632E9A8A0BDD
     created: 2020-11-13  expires: never       usage: SC
     trust: unknown       validity: unknown
ssb  rsa4096/DBCD8B98F2F9188C
     created: 2020-11-13  expires: never       usage: E
[ unknown] (1). Ponder Stibbons <ponder.stibbons@unseen.edu>
Type 5 to trust your keys completely.

Please decide how far you trust this user to correctly verify other users' keys
(by looking at passports, checking fingerprints from different sources, etc.)

  1 = I don't know or won't say
  2 = I do NOT trust
  3 = I trust marginally
  4 = I trust fully
  5 = I trust ultimately
  m = back to the main menu

Your decision? 5
Confirm your choice by entering Y.

Do you really want to set this key to ultimate trust? (y/N) y

sec  rsa4096/C8DE632E9A8A0BDD
     created: 2020-11-13  expires: never       usage: SC
     trust: ultimate      validity: unknown
ssb  rsa4096/DBCD8B98F2F9188C
     created: 2020-11-13  expires: never       usage: E
[ unknown] (1). Ponder Stibbons <ponder.stibbons@unseen.edu>
Please note that the shown key validity is not necessarily correct
unless you restart the program.
Use the command quit to exit.

gpg> quit
