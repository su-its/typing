import React from "react";
import { Avatar, Box, Text, HStack, VStack } from "@chakra-ui/react";
import { getCurrentUser } from "@/app/actions";
import type { User } from "@/types/user";

interface UserCardPresenterProps {
  user?: User;
}

export const UserCardPresenter = ({ user }: UserCardPresenterProps) => {
  return (
    <Box bg="blue.600" p={5}>
      <HStack spacing={4}>
        <Avatar /*src={ TODO: しずっぴーを表示 }*/ maxW="100px" borderRadius="9" />
        <VStack align="start">
          <Text fontSize="lg" fontWeight="bold" color="white">
            名前: {user ? user.handleName : "ログインしていません"}
          </Text>
          <Text color="white">学籍番号: {user ? user.studentNumber : "未ログイン"}</Text>
        </VStack>
      </HStack>
    </Box>
  );
};

const UserCard = async () => {
  const user = await getCurrentUser();
  return <UserCardPresenter user={user} />;
};

export default UserCard;
