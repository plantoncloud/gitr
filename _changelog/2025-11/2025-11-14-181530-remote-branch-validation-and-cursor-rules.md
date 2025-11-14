# Remote Branch Validation with Fallback & Cursor Rules Integration

**Date**: November 14, 2025

## Summary

Enhanced the `gitr rem` command to validate branch existence on remote before opening in browser, with intelligent fallback to the default branch when the local branch hasn't been pushed yet. Additionally, established a complete cursor rules framework for gitr with commit conventions, PR generation, and changelog documentation rules similar to planton-cloud.

## Problem Statement

Users running `gitr rem` on local-only branches (branches that haven't been pushed to the remote) were experiencing a frustrating workflow issue: the command would successfully open their browser, but the web page would display a "404 - Page not found" error because the branch doesn't exist on GitHub/GitLab/Bitbucket yet.

### Pain Points

- **Poor User Experience**: Browser opens to a non-existent URL, requiring manual navigation to find the correct page
- **Confusing Behavior**: No warning or indication that the branch doesn't exist remotely
- **Workflow Interruption**: Users had to manually determine the default branch and construct the correct URL
- **SSH Authentication Failures**: Initial implementation attempts using SSH connections would fail with authentication errors, blocking the entire command
- **Missing Development Infrastructure**: No standardized workflow for commits, PR creation, or documentation in the gitr repository

## Solution

Implemented a two-pronged approach:

### 1. Local Remote-Tracking Branch Validation

Instead of querying the remote repository via SSH (which requires authentication and network access), the solution checks **local remote-tracking branches** stored in `.git/refs/remotes/origin/`. This approach:

- Works offline (uses only local git data)
- Requires no authentication
- Executes instantly (no network latency)
- Provides reliable information from the last `git fetch`/`git pull`

### 2. Intelligent Default Branch Fallback

When a local branch doesn't exist on the remote:

- Display a clear warning message to the user
- Automatically detect the repository's default branch by checking `refs/remotes/origin/HEAD`
- Fall back to common defaults (`main`, `master`) if HEAD reference is not found
- Open the default branch instead of the non-existent branch

### 3. Cursor Rules Framework

Established gitr-specific cursor rules modeled after planton-cloud patterns:

- **Commit rule**: Conventional commit message generation with gitr-aware scopes
- **PR info rule**: Automated PR title and description generation
- **Changelog rule**: Structured documentation for significant changes

### Key Components

**New Functions in `pkg/git/git.go`**:
- `DoesBranchExistOnRemote()`: Checks local remote-tracking branches without network access
- `GetDefaultBranch()`: Determines repository default branch from local references

**Updated Command Handler in `cmd/gitr/root/web.go`**:
- Branch validation logic integrated into `rem` command case
- Warning messages with clear user feedback
- Graceful fallback handling

**Cursor Rules**:
- `.cursor/rules/git/commit-gitr-changes.mdc`
- `.cursor/rules/git/github/pull-requests/generate-gitr-pr-info.mdc`
- `.cursor/rules/git/create-gitr-changelog.mdc`

## Implementation Details

### Remote Branch Validation

```go
// DoesBranchExistOnRemote checks if a branch exists on the remote repository
// by checking local remote-tracking branches (e.g., refs/remotes/origin/branch-name)
func DoesBranchExistOnRemote(r *git.Repository, branchName string) bool {
    remotes, err := r.Remotes()
    if err != nil || len(remotes) == 0 {
        return false
    }

    remoteName := remotes[0].Config().Name
    remoteTrackingRef := "refs/remotes/" + remoteName + "/" + branchName
    
    refs, err := r.References()
    if err != nil {
        return false
    }

    exists := false
    refs.ForEach(func(ref *plumbing.Reference) error {
        if ref.Name().String() == remoteTrackingRef {
            exists = true
            return errors.New("found") // stop iteration
        }
        return nil
    })

    return exists
}
```

### Default Branch Detection

```go
// GetDefaultBranch returns the default branch of the remote repository
// by checking the local remote HEAD reference (e.g., refs/remotes/origin/HEAD)
func GetDefaultBranch(r *git.Repository) (string, error) {
    remotes, err := r.Remotes()
    if err != nil || len(remotes) == 0 {
        return "", errors.New("no remotes found")
    }

    remoteName := remotes[0].Config().Name
    remoteHeadRef := "refs/remotes/" + remoteName + "/HEAD"
    
    ref, err := r.Reference(plumbing.ReferenceName(remoteHeadRef), true)
    if err == nil {
        targetRef := ref.Name().String()
        if ref.Type() == plumbing.SymbolicReference {
            targetRef = ref.Target().String()
        }
        
        defaultBranch := strings.TrimPrefix(targetRef, "refs/remotes/"+remoteName+"/")
        if defaultBranch != "" && defaultBranch != targetRef {
            return defaultBranch, nil
        }
    }

    // Fallback: try common default branch names
    commonDefaults := []string{"main", "master"}
    for _, defaultBranch := range commonDefaults {
        if DoesBranchExistOnRemote(r, defaultBranch) {
            return defaultBranch, nil
        }
    }

    return "", errors.New("unable to determine default branch")
}
```

### Command Handler Integration

```go
case rem:
    branchToOpen := branch
    // Check if the current branch exists on the remote
    if !git.DoesBranchExistOnRemote(r, branch) {
        log.Warnf("Branch '%s' doesn't exist on remote. Opening default branch instead.", branch)
        defaultBranch, err := git.GetDefaultBranch(r)
        if err != nil {
            log.Warnf("Unable to determine default branch: %v. Attempting to open '%s' anyway.", err, branch)
        } else {
            branchToOpen = defaultBranch
        }
    }
    url.OpenInBrowser(web.GetRemUrl(s.Provider, webUrl, branchToOpen))
```

### Build System Improvements

- **Upgraded GoReleaser to v2**: Updated `.goreleaser.yaml` with `version: 2` directive
- **Fixed deprecated syntax**: Changed `--skip-publish` to `--skip=publish`
- **Updated Makefile**: Changed build directory from `build` to `dist`, updated path patterns for v2 output structure
- **Modernized brews section**: Changed `tap:` to `repository:`, `folder:` to `directory:`

### Cursor Rules Structure

Created comprehensive rules following planton-cloud patterns:

**Commit Rule** (`commit-gitr-changes.mdc`):
- Conventional Commits format (feat, fix, refactor, etc.)
- Gitr-specific scopes: cmd, pkg/git, pkg/web, pkg/config, site, build, docs
- Example commit messages for various scenarios
- Execution guidelines and error handling

**PR Info Rule** (`generate-gitr-pr-info.mdc`):
- Generates PR title and description as two copyable code blocks
- Component-aware scope detection
- Structured sections: Summary, Context, Changes, Implementation notes, Breaking changes, Test plan, Risks, Checklist
- Example outputs demonstrating the format

**Changelog Rule** (`create-gitr-changelog.mdc`):
- Proportional sizing guidance (small: 150-300, medium: 300-600, large: 600-1000+ lines)
- Required sections and optional sections
- Gitr-specific guidelines for commands, providers, workflows
- Quality checklist and anti-patterns
- Timestamp-based file naming with year-month organization

## Benefits

### User Experience

- **No more 404 errors**: Users always reach a valid page when using `gitr rem`
- **Clear feedback**: Warning messages inform users when fallback occurs
- **Seamless workflow**: Automatic fallback means no manual URL construction needed
- **Works offline**: Validation uses local data, no network required

### Performance

- **Instant validation**: No network latency or SSH handshake delays
- **Reduced error surface**: No authentication failures blocking functionality
- **Efficient**: Reads local git references without spawning external processes

### Developer Experience

- **Reliable behavior**: Consistent functionality regardless of SSH configuration
- **Better error handling**: Graceful degradation with informative messages
- **Standardized workflow**: Cursor rules provide consistent patterns for commits, PRs, and documentation

### Code Quality

- **Clean separation of concerns**: Branch validation logic isolated in dedicated functions
- **Testable code**: Pure functions that can be unit tested
- **Maintainable**: Clear function names and documentation
- **Consistent conventions**: Cursor rules enforce commit message and PR description standards

## Impact

### Users

- Eliminates frustrating 404 errors when working with local branches
- Saves time by automatically navigating to a valid page
- Provides clear information about what's happening via warning messages
- Works reliably even with SSH authentication issues

### Developers

- Reduces support burden (fewer user complaints about "broken" web commands)
- Establishes clear conventions for commits and PRs
- Provides structured approach to documentation
- Aligns gitr development patterns with planton-cloud

### Workflows

- **Feature branches**: Developers can use `gitr rem` before pushing, fallback to main
- **Code review**: Better visibility with structured PR descriptions
- **Knowledge preservation**: Changelogs capture rationale and design decisions
- **Onboarding**: New contributors can follow established patterns via cursor rules

### Supported Providers

All SCM providers continue to work correctly:
- ✅ **GitHub**: `https://github.com/owner/repo/tree/branch`
- ✅ **GitLab**: `https://gitlab.com/owner/repo/-/tree/branch`
- ✅ **Bitbucket Cloud**: `https://bitbucket.org/owner/repo/branch/branch`
- ✅ **Bitbucket Datacenter**: Self-hosted instances

## Testing Strategy

### Manual Testing

1. **Local-only branch scenario**:
   ```bash
   git checkout -b test-local-branch
   gitr rem
   # Expected: Warning message + opens default branch
   ```

2. **Pushed branch scenario**:
   ```bash
   git checkout main
   gitr rem
   # Expected: No warning + opens main branch
   ```

3. **Build verification**:
   ```bash
   make local
   # Expected: Successful build with GoReleaser v2
   ```

### Test Results

- ✅ Warning message displays correctly for unpushed branches
- ✅ Default branch detection works (falls back to main)
- ✅ No SSH authentication errors
- ✅ Browser opens to correct URL
- ✅ Existing pushed branches still work normally
- ✅ GoReleaser v2 build succeeds
- ✅ Binary installs correctly to ~/bin/gitr

## Breaking Changes

**None** - This is a backward-compatible enhancement:

- Existing functionality preserved for pushed branches
- New validation only adds safety for edge cases
- Warning messages are informational only
- Users can still override by pushing branches first

## Known Limitations

1. **Stale remote-tracking branches**: If a user hasn't run `git fetch` recently, the local remote-tracking branches may be out of sync. The validation reflects what's known locally, not the current remote state.

2. **Remote HEAD configuration**: Some repositories may not have `refs/remotes/origin/HEAD` configured. The fallback to common defaults (main, master) handles this, but unusual default branch names may not be detected.

3. **Multiple remotes**: The implementation uses the first remote found. Repositories with multiple remotes may not select the user's intended remote.

## Future Enhancements

- **Cache remote branch list**: Periodically fetch and cache remote branch information
- **Prompt for action**: Offer to push the branch or select from available branches
- **Multi-remote support**: Allow users to specify which remote to use
- **Branch creation workflow**: Integrate with GitHub/GitLab APIs to create branches remotely
- **Dry-run improvements**: Show which branch would be opened without actually opening it

## Related Work

This change complements:

- **Existing web commands**: `branches`, `commits`, `prs`, `issues`, `tags`, `releases`, `pipelines`
- **Clone functionality**: Maintains consistency with gitr's SCM provider support
- **Configuration system**: Uses existing config to determine provider and hostname
- **Future PR workflow**: The new cursor rules will streamline contribution process

## Code Metrics

### Files Changed
- `pkg/git/git.go`: +74 lines (2 new functions)
- `cmd/gitr/root/web.go`: +10 lines (validation logic)
- `.goreleaser.yaml`: Modified for v2 compatibility
- `Makefile`: Updated for dist directory and path patterns
- `.gitignore`: Added dist and .cursor/workspace/
- `.cursor/rules/`: +3 new rule files (~800 lines total)
- `_changelog/`: New directory with README

### Dependencies
- No new external dependencies
- Uses existing `go-git/go-git/v5` and `go-git/go-git/v5/plumbing`

## Design Decisions

### Why Local Remote-Tracking Branches?

**Considered alternatives**:
1. ❌ **SSH query to remote**: Requires authentication, fails with SSH issues, slow
2. ❌ **HTTP API calls**: Requires tokens, rate limiting, different per provider
3. ✅ **Local remote-tracking branches**: Fast, reliable, no auth needed

**Trade-off**: Accuracy depends on how recently `git fetch` was run, but this is an acceptable trade-off for the reliability and performance benefits.

### Why Automatic Fallback vs. User Prompt?

**Considered alternatives**:
1. ❌ **Error and abort**: Frustrating UX, forces manual navigation
2. ❌ **Interactive prompt**: Breaks non-interactive usage, requires user input
3. ✅ **Automatic fallback with warning**: Best of both worlds - informative and seamless

**Rationale**: Users running `gitr rem` want to quickly view their repository in the browser. An automatic fallback with a clear warning message achieves this goal while informing the user about what happened.

### Why Cursor Rules?

**Rationale**: Establishing cursor rules now creates a foundation for consistent development practices as gitr grows. The rules are modeled after planton-cloud's proven patterns, bringing best practices to this repository.

---

**Status**: ✅ Production Ready

**Timeline**: ~4 hours of development, testing, and documentation

**Contributors**: Suresh (with AI assistance)

