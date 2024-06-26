import { Container, Stack } from '@chakra-ui/react';
import Navbar from './components/Navbar';
import TodoForm from './components/TodoForm';
import TodoList from './components/TodoList';

// Only use this line in production:
// export const BASE_URL = import.meta.env.MODE === "development" ? "http://127.0.0.1:5000/api" : "/api"

// Use this line in development:
export const BASE_URL = "http://127.0.0.1:5000/api"

function App() {
  return (
    <Stack h="100vh">
      <Navbar />
      <Container>
        <TodoForm />
        <TodoList />
      </Container>
      {/* <Button>Hello</Button> */}
    </Stack>
  )
}

export default App
