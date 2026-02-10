# Deploying Go Maps App to Vercel

1. Install the Vercel CLI and login:

```bash
npm i -g vercel
vercel login
```

2. From the project root, deploy (follow prompts):

```bash
vercel --prod
```

Vercel will detect the `api/` function and use `vercel.json` to route `/map` and `/` to the function.

Postman usage

1. Import `postman_collection.json` into Postman.
2. Set the `base_url` variable to your deployment URL (for example `https://your-app.vercel.app`).
3. Use the `Map Redirect` request with `location` query parameter to test. Example:

```
GET https://your-app.vercel.app/map?location=Paris
```
