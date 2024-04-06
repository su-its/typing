import React from "react";
import { Avatar, Box, Text, HStack, VStack } from "@chakra-ui/react";
import { getCurrentUser } from "@/app/actions";
import type { User } from "@/types/user";

interface UserCardPresenterProps {
  user?: User;
}

export const UserCardPresenter = ({ user }: UserCardPresenterProps) => {
  return (
    <HStack spacing={4} bg="blue.600">
      <Avatar
        src={"https://www.shizuoka.ac.jp/cms/files/shizudai/MASTER/0100/uISrbYCb_VL033_r03.png"}
        boxSize="100px"
        borderRadius="0"
      />
      <VStack align="start" paddingRight="8px">
        <Text fontSize="lg" fontWeight="bold" color="white">
          名前: {user ? user.handleName : "ログインしていません"}
        </Text>
        <Text color="white">学籍番号: {user ? user.studentNumber : "未ログイン"}</Text>
      </VStack>
    </HStack>
  );
};

const UserCard = async () => {
  const user = await getCurrentUser();
  return <UserCardPresenter user={user} />;
};

export default UserCard;
