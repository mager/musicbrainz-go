.PHONY: tag-patch publish

# Get the latest tag and increment the patch version
tag-patch:
	@latest_tag=$$(git describe --tags --abbrev=0 2>/dev/null || echo "v0.0.0"); \
	major=$$(echo $$latest_tag | cut -d. -f1 | tr -d 'v'); \
	minor=$$(echo $$latest_tag | cut -d. -f2); \
	patch=$$(echo $$latest_tag | cut -d. -f3); \
	new_patch=$$((patch + 1)); \
	new_tag="v$$major.$$minor.$$new_patch"; \
	echo "Creating new tag: $$new_tag"; \
	git tag $$new_tag; \
	echo "Tag $$new_tag created. Don't forget to push it with: git push origin $$new_tag"

# Create a new tag and push it to remote
publish: tag-patch
	@latest_tag=$$(git describe --tags --abbrev=0); \
	echo "Pushing tag $$latest_tag to remote..."; \
	git push origin $$latest_tag; \
	echo "Tag $$latest_tag published successfully!"
