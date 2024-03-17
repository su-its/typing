import React from "react";
import { Avatar, Box, Text, HStack, VStack } from "@chakra-ui/react";

const UserCard: React.FC = () => {
  const bgColor = "#0000cd"; // 濃い青の背景色

  // モックのユーザー情報
  const user = {
    name: "テストユーザー",
    studentId: "24AB1234",
    avatarUrl: "https://avatars.githubusercontent.com/u/12345678?v=4",
  };

  return (
    user && (
      <Box bg={bgColor} p={4} maxH="120px">
        <HStack spacing={4}>
          <Avatar src={user.avatarUrl} maxW="100px" borderRadius="9" />
          <VStack align="start" spacing={1}>
            <Text fontSize="lg" fontWeight="bold">
              名前: {user.name}
            </Text>
            <Text>学籍番号: {user.studentId}</Text>
          </VStack>
        </HStack>
      </Box>
    )
  );
};

export default UserCard;
