# Skill: Deploy to Cloud Run

## Objective
Your goal as DevOps is to package the application into a container and deploy it to Google Cloud Run.

## Instructions
1. **Verify Environment**: Ensure the necessary files for the chosen tech stack are in `src/`.
2. **Containerize**: Use the IDE terminal to navigate to `src/` and run `gcloud run deploy --source .`. 
3. **Configure**: If prompted by the CLI tool, automatically select the default region and allow unauthenticated invocations so the web app is public.
4. **Report**: Output the live production Google Cloud Run URL to the user!
5. **CI/CD**: Check and follow the instructions in the `.agents/rules/github-actions-for-Flutter.md` file before taking any actions regarding CI/CD
