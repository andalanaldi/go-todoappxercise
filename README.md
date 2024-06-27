# Fullstack React and Go ToDo App 

I learn Go by practice with a fullstack project that I follow through tutorials on youtube. The tutorials are from freeCodeCamp and Burak Orkmez as a tutor from As a Programmer YouTube Channel. Hand-on experience in full stack development like this is important to get a better understanding of a fullstack react and golang. For those who are curious, you can try this practice/exercise. Here are the Resources for your reference:

## Useful Resources

- Check out this [freeCodeCamp video](https://www.youtube.com/watch?v=lNd7XlXwlho) for a comprehensive guide!
- Watch the insightful [As a Programmer video](https://www.youtube.com/watch?v=zw8z_o_kDqc) on YouTube!
- Visit the [As a Programmer channel](https://www.youtube.com/@asaprogrammer_) for more content!
- Explore Burak Orkmez's [React Go tutorial](https://github.com/burakorkmez/react-go-tutorial) on GitHub!
- Connect with [Aldi Andalan on LinkedIn](https://www.linkedin.com/in/aldi-andalan/)!

It is still ongoing so, I will update again later on.

## Tech Stack

### Backend
- Golang (Go Programming Language)
- Air (Golang equivalent of nodemon in Javascript)
- Fiber (Golang library for API)
- Bson (Golang library package for binary JSON)
- MongoDB (NoSQL Database to store Todo list related data)

### Frontend
- React (Javascript library to build Frontend app)
- Chakra UI (UI library in react to give style and CSS for ToDo App)
- React icons (imported to display icons in Navbar)
- Tan Stack (To fetch API and query, NOT BEING USED YET)

Deployment or Portfolio Documentation ? (will be updated soon)

### Set up on local development environment

This repo could be cloned and run locally to check whether the Todo App is running in local. The database is available on Author's MongoDB Atlas account. So, Internet connection is enough to run this Todo app locally.
First things first we can set up Golang first (Do not forget to install [Golang](https://go.dev/doc/install), [Node.JS](https://nodejs.org/en) and [Typescript](https://www.npmjs.com/package/typescript) with its set policy on Windows Locally, please try to [install Typescript](https://stackoverflow.com/questions/35369501/tsc-is-not-recognized-as-internal-or-external-command) on CMD rather than on PWSH in VS Code terminal to make sure it is recognized in local).  

- First do a similar command with npm init in go :
```Golang
go mod init github.com/<your repo>
```

- This is for checking whether Go is installed and which verison it is
```Golang
go version
```

- Similar to npm run start
```Golang
go run main.go
```

- Install an equivalent to nodemon
```Golang
go install github.com/air-verse/air@latest
```
please keep in mind that cosmtrek is replaced by air-verse for the latest update

- Install all the necessary library from Golang
```Golang
go get github.com/gofiber/fiber/v2

go get github.com/joho/godotenv

go get go.mongodb.org/mongo-driver/mongo

go.mongodb.org/mongo-driver/bson
```

- Install all the necessary library for React
```Typescript
npm create vite@latest

npm install

npm run dev

npm i @chakra-ui/react @emotion/react @emotion/styled framer-motion

npm install react-icons

npm i @tanstack/react-query

npm run build
```

Visit this app's backend documentation [here](https://documenter.getpostman.com/view/29042682/2sA3dsnESF)!