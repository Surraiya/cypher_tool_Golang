# Cypher Tool

Cypher tool is a Command-Line Encryption and Decryption Tool. This command-line tool allows users to encrypt and decrypt messages using various encryption techniques. It supports three encryption methods: ROT13, Reverse, and Caesar Cipher.

## Tool Description

This tool provides a user-friendly interface to perform encryption and decryption of messages. It starts with a greeting and then guides the user through the following steps:

1. Select the operation (Encrypt or Decrypt).
2. Choose the encryption type (ROT13, Reverse, or Caesar Cipher).
3. Input the message to encrypt or decrypt.
4. Output the result of the chosen operation.

The tool also handles user input validation, ensuring that non-alphabet characters in the message are left unchanged.

## Tool Usage

To use the tool, follow these steps:

1. Open your command-line terminal.

2. Run the tool by executing the script or binary (assuming the name of the tool's executable is encrypt_tool):
<pre>
```bash
Copy code
./encrypt_tool
```
</pre>

3. You will be greeted and prompted to select the operation (Encrypt or Decrypt), like so:
<pre>
```bash
Welcome to the Cypher Tool!

Select operation (1/2):
1. Encrypt.
2. Decrypt.
```
</pre>

4. Choose an operation by entering the corresponding number (1 for Encrypt or 2 for Decrypt).

5. Next, select the encryption type. You will be prompted with options like:
<pre>
```bash
Select cypher (1/3):
1. ROT13.
2. Reverse.
3. Caesar
```
</pre>

6. Choose an encryption method by entering the corresponding number (1 for ROT13, 2 for Reverse, or 3 for Caesar Cipher).

7. Input the message you want to encrypt or decrypt:
<pre>
```bash
Enter the message:
```
</pre>

8. The tool will perform the selected operation and display the result:
<pre>
```bash
Result:
[encrypted/decrypted message]
```
</pre>
9. If you want to perform another encryption or decryption, you can restart the tool or exit it.
```bash
Do you want to perform another operation? (y/n)
```
</pre>


## Encryption Methods
The tool supports the following encryption techniques:

### ROT13: 
ROT13 is a simple letter substitution cipher that replaces a letter with the 13th letter after it in the alphabet. It is a symmetric encryption method, meaning that applying ROT13 twice will result in the original text.

<b> Example: </b>
<pre>
```bash
Enter a message: 
Hello World!
Encrypted Message using ROT13: 
"Uryyb, Jbeyq!"
Decrypted Message using ROT13: 
"Hello, World!"
```
</pre>

### Reverse: 
The Reverse encryption method reverses the order of the characters in the message. This method is straightforward and symmetric.

<b> Example: </b>
<pre>
```bash
Enter a message: 
Hello World!
Encrypted message using Reverse: 
"Svool Dliow!"
Decrypted message using Reverse:
Hello World!
```
</pre>

### Caesar: 
The Caesar Cipher allows the user to specify a custom shift value (key) for encryption or decryption. It shifts each letter in the message by the specified key, wrapping around the alphabet if necessary. The Caesar Cipher is also symmetric, meaning that the same key is used for both encryption and decryption.

<b> Example: </b>
<pre>
```bash
Enter a message: 
Hello World!
Enter the Caesar Cipher shift you want: 
4
Encrypted message using Caesar technique:
Lipps Asvph!
Decrypted message using Caesar technique:
Hello World!
```
</pre>

Feel free to use this command-line tool to encrypt or decrypt messages with ease, whether for fun or secure communication!

