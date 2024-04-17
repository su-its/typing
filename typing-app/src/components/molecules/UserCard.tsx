import React from "react";
import { Avatar, Text, HStack, VStack, Spacer } from "@chakra-ui/react";
import type { StackProps } from "@chakra-ui/react";
import { getCurrentUser } from "@/app/actions";
import type { User } from "@/types/user";

interface UserCardPresenterProps extends StackProps {
  user?: User;
}

export const UserCardPresenter = ({ user, ...rest }: UserCardPresenterProps) => {
  const props: StackProps = {
    width: rest?.width ?? "18%",
    ...rest,
  };

  return (
    <HStack spacing={4} bg="blue.600" {...props}>
      <Avatar
        src={"https://www.shizuoka.ac.jp/cms/files/shizudai/MASTER/0100/uISrbYCb_VL033_r03.png"}
        boxSize="70px"
        borderRadius="0"
      />
      <VStack align="start" overflow="hidden" flexGrow={1} gap={0}>
        <Text fontSize="lg" fontWeight="bold" color="white" isTruncated width="90%">
          名前: {user ? user.handleName : "ログインしていません"}
        </Text>
        <Text color="white" width="90%" isTruncated>
          学籍番号: {user ? user.studentNumber : "未ログイン"}
        </Text>
      </VStack>
    </HStack>
  );
};

const UserCard = async (props?: StackProps) => {
  const user = await getCurrentUser();
  return <UserCardPresenter user={user} {...props} />;
};

export default UserCard;
