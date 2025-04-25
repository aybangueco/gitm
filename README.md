# Gitm

A simple cli application to manage multiple git cli accounts.

## Motivation

I noticed that there is no built in functionality in git itself, when managing accounts especially, if you have multiple personality (account). Typing git config commands is kind of a hassle repeatedly, that's why i created this simple cli application.

## Installation

```bash
go install github.com/aybangueco/gitm
```

## Commands

### Init

This command is mandatory, especially for fresh installs of Gitm.

```bash
gitm init
```

### Account list

Shows list of accounts.

```bash
gitm list-accounts
```
### Switch Account
Switch to an existing accounts (replace account id with actual id of account).
```bash
gitm switch-account [accountID]
```

### Add Account

Adds new git account.

```bash
gitm add-account
```

### Update Account

Update information of existing accounts (replace account id with actual id of account).

```bash
gitm update-account [accountID]
```

### Delete Account

Delete existing account (replace account id with actual id of account).

```bash
gitm delete-account [accountID]
```

## TODO

- [x] Add account
- [x] Display lists of account in a table
- [x] Update account (pending)
- [x] Switch account
- [x] Delete account
- [ ] Add ssh key specifics on account

## Contributing

1. Fork this repository
2. Clone your forked repository:
3. Create a new branch

```bash
git checkout -b feature/your-feature-name
```

4. Make your changes and commit

```bash
git add .
git commit -m "feat: add your feature description"
```

5. Push your fork
6. Create a pull request (pr) to the main repo
7. Wait for a review and make any requested change
