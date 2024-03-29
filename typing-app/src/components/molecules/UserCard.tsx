import React from "react";
import { Avatar, Box, Text, HStack, VStack } from "@chakra-ui/react";

const UserCard: React.FC = () => {

  // モックのユーザー情報
  const user = {
    name: "テストユーザー",
    studentId: "24AB1234",
    avatarUrl: "https://avatars.githubusercontent.com/u/12345678?v=4",
  };

  return (
    user && (
      <Box bg={ "blue.600" } p={5} >
        <HStack spacing={4}>
          <Avatar src={user.avatarUrl} maxW="100px" borderRadius="9" />
          <VStack align="start">
            <Text fontSize="lg" fontWeight="bold" color={"white"}>
              名前: {user.name}
            </Text>
            <Text color={"white"}>学籍番号: {user.studentId}</Text>
          </VStack>
        </HStack>
      </Box>
    )
  );
};

export default UserCard;
