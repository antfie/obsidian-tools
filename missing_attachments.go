package main

func (ctx *Context) FindMissingAttachments(source string) error {
	sourceAbs := assertSourceExists(source)

	sourceObsidianRoot, err := findObsidianRoot(sourceAbs)

	if err != nil {
		return err
	}

	// This will output any warnings about missing attachments
	err = ctx.Repository.PopulateFromVault(sourceObsidianRoot, true)

	return err
}
