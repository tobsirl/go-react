import { Button } from "@/components/ui/button";
import { Stack } from "@chakra-ui/react";

function App() {
  return (
    <>
      <Stack h="100vh">
        <Button color={"white"} bg={"blue.500"} variant={"outline"}>
          Hello
        </Button>
        <Button>Click me</Button>
      </Stack>
    </>
  );
}

export default App;
