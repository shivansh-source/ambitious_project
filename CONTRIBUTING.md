# Contributing to VideEdit

Thank you for considering contributing to VideEdit! This document provides guidelines and instructions for contributing.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for all contributors.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check existing issues. When creating a bug report, include:

- **Clear title and description**
- **Steps to reproduce** the behavior
- **Expected vs actual behavior**
- **Screenshots** if applicable
- **Environment details** (OS, browser, versions)

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion, include:

- **Clear title and description**
- **Use cases** for the enhancement
- **Potential implementation** approach if you have ideas
- **Mockups or examples** if applicable

### Pull Requests

1. **Fork the repository** and create your branch from `main`
2. **Make your changes** following our coding standards
3. **Add tests** for new functionality
4. **Ensure tests pass** locally
5. **Update documentation** as needed
6. **Write clear commit messages**
7. **Submit the pull request**

## Development Setup

See [Developer Guide](docs/developer-guide/GETTING_STARTED.md) for complete setup instructions.

## Coding Standards

### JavaScript/TypeScript
- Use ESLint and Prettier
- Follow Airbnb style guide
- Write meaningful variable names
- Add JSDoc comments for functions

### Python
- Follow PEP 8
- Use Black for formatting
- Use type hints
- Add docstrings for functions

### Git Commits
Follow conventional commits:
- `feat:` new feature
- `fix:` bug fix
- `docs:` documentation changes
- `style:` formatting changes
- `refactor:` code refactoring
- `test:` adding tests
- `chore:` maintenance tasks

Example: `feat: add scene detection AI service`

## Testing

- Write unit tests for new functionality
- Ensure all tests pass before submitting PR
- Aim for >80% code coverage
- Test edge cases and error conditions

## Documentation

- Update README if adding features
- Add API documentation for new endpoints
- Include code comments for complex logic
- Update architecture docs for significant changes

## Project Structure Guidelines

### Adding New Features

When adding a new feature:

1. **Backend API**: Add route in `backend/api/routes/`
2. **Service Logic**: Implement in `backend/services/`
3. **Frontend Component**: Add to `frontend/src/components/`
4. **AI Service**: Create in `ai-services/`
5. **Tests**: Add tests in appropriate `tests/` directory
6. **Documentation**: Update relevant docs

### File Organization

- Keep files focused and single-purpose
- Use descriptive file names
- Group related functionality
- Avoid deep nesting (max 4 levels)

## Review Process

1. **Automated checks** must pass (linting, tests)
2. **Code review** by maintainer
3. **Testing** of functionality
4. **Documentation review**
5. **Merge** after approval

## Questions?

Feel free to ask questions in:
- GitHub Discussions
- Issue comments
- Pull request comments

Thank you for contributing! 🎉
