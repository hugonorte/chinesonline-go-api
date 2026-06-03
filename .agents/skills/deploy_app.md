# Skill: Deploy App

## Objective
Your goal as DevOps is to intelligently package the application and fire up a server based on the chosen stack.

## Instructions
1. **Stack Detection**: Inspect the `Technical_Specification.md` and the files in `src/` to figure out what stack is being used.

2. **Install Dependencies**: Always ask for user permission before installing any npm dependency. If user agree to proceed with a new dependency installation, use your native terminal to navigate into `./` and run `npm install`, 

3. **Host Locally**: Execute the appropriate native terminal command (e.g., `flutter run`) to start a background server.
4. **Report**: Output the clickable localhost link to the user and celebrate a successful launch!
5. **CI/CD**: Check and follow the instructions in the `.agents/rules/github-actions-for-Flutter.md` file before taking any actions regarding CI/CD
