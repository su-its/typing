// components/MenuContainer.js
import { Flex, VStack } from "@chakra-ui/react";
import GameStartButton from "../atoms/gameStartButton";
import RankingButton from "../atoms/rankingButton";
import LogoutButton from "../atoms/logoutButton";

const HomeMenuContainer = () => {
    return (
        <Flex justify="center" align="center" h="65vh">
            <VStack spacing = {8} align="stretch" mx="auto">
                <GameStartButton />
                <RankingButton />
                <LogoutButton />
            </VStack>
        </Flex>
    );
};

export default HomeMenuContainer;
