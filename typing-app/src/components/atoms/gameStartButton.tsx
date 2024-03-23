// components/buttons/GameStartButton.js
'use client';

import Link from 'next/link';
import { Button } from "@chakra-ui/react";

const GameStartButton = () => {
    return (
        <Button as={Link} href="game" colorScheme="green" size="lg">
            Game Start
        </Button>
    );
};

export default GameStartButton;
