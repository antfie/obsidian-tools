package main

import "errors"

var (
	ErrSourceVaultNotResolved                  = errors.New("the source vault could not be resolved")
	ErrDestinationVaultNotResolved             = errors.New("the destination vault could not be resolved")
	ErrDestinationIsSameVaultAsSource          = errors.New("the destination vault is the same as the source vault. Use Obsidian for vault-based file operations")
	ErrNotMarkdownFile                         = errors.New("this operation only works on markdown files")
	ErrCouldNotAssertDestinationFilePathUnique = errors.New("could not assert the destination path was unique")
)
