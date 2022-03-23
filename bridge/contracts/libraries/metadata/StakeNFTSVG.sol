// SPDX-License-Identifier: GPL-2.0-or-later
pragma solidity >=0.7.6;

import "@openzeppelin/contracts/utils/Strings.sol";
import "contracts/utils/Base64.sol";

/// @title StakeNFTSVG
/// @notice Provides a function for generating an SVG associated with a Staked NFT position
library StakeNFTSVG {
    using Strings for uint256;

    struct StakeNFTSVGParams {
        string shares;
        string freeAfter;
        string withdrawFreeAfter;
        string accumulatorEth;
        string accumulatorToken;
    }

    function generateSVG(StakeNFTSVGParams memory params)
        internal
        pure
        returns (string memory svg)
    {
        return string(abi.encodePacked(generateSVGDefs(params), generateSVGText(params), "</svg>"));
    }

    function generateSVGText(StakeNFTSVGParams memory params)
        private
        pure
        returns (string memory svg)
    {
        svg = string(
            abi.encodePacked(
                "<text x='10' y='20'>Shares: ",
                params.shares,
                "</text>",
                "<text x='10' y='40'>Free after: ",
                params.freeAfter,
                "</text>",
                "<text x='10' y='60'>Withdraw Free After: ",
                params.withdrawFreeAfter,
                "</text>",
                "<text x='10' y='80'>Accumulator (ETH): ",
                params.accumulatorEth,
                "</text>",
                "<text x='10' y='100'>Accumulator (Token): ",
                params.accumulatorToken,
                "</text>"
            )
        );
    }

    function generateSVGDefs(StakeNFTSVGParams memory params)
        private
        pure
        returns (string memory svg)
    {
        svg = string(
            abi.encodePacked(
                '<svg width="500" height="500" viewBox="0 0 500 500" xmlns="http://www.w3.org/2000/svg"',
                " xmlns:xlink='http://www.w3.org/1999/xlink'>"
            )
        );
    }
}
