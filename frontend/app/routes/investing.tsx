import { Button, Text } from "@mantine/core";
import { Outlet } from "@remix-run/react";

export default function Index() {
  return (
    <div style={{ fontFamily: "system-ui, sans-serif", lineHeight: "1.4" }}>
      <Text>Hello world!</Text>
    </div>
  );
}
