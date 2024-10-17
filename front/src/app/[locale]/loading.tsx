"use client";

import Lottie from "react-lottie-player/dist/LottiePlayerLight";

import lottieJson from "../../../public/loader.json";
import styled from "styled-components";

export default function Loading() {
  return (
    <LoaderWrapper>
      <Lottie
        loop
        animationData={lottieJson}
        play
        style={{ width: 350, height: 350 }}
      />
    </LoaderWrapper>
  );
}

const LoaderWrapper = styled.div`
  display: grid;
  place-content: center;
  height: 100vh;
  opacity: 0.6;
`;
