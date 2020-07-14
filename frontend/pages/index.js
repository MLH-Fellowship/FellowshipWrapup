// Libraries
import { useState, useEffect } from "react";
import Head from "next/head";
import dynamic from "next/dynamic";
import Cookie from "js-cookie";
import { parseCookies } from "../utils/parseCookies";
import useWindowSize from "react-use/lib/useWindowSize";
import Confetti from "react-confetti";

// Components
import Header from "../components/Landing/Header";
import Example from "../components/Landing/Example";

export default function Home({ cookies }) {
  const [isFirstVisit, setIsFirstVisit] = useState(
    cookies.firstVisit ? cookies.firstVisit : true
  );

  useEffect(() => {
    if (cookies.firstVisit === undefined) {
      Cookie.set("firstVisit", true);
      setIsFirstVisit(true);
    }

    if (cookies.firstVisit === "true") {
      Cookie.set("firstVisit", false);
      setIsFirstVisit(false);
    }
  }, [cookies]);

  return (
    <div>
      <Head>
        <title>Fellowship Wrapup</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main>
        {/* I had to use hard coded values for now, since you can't access the window object with server-side rendering and couldn't find an easy and fast alternative */}
        {isFirstVisit && (
          <Confetti
            width={1500}
            height={1000}
            recycle={false}
            opacity={0.8}
            numberOfPieces={800}
          />
        )}

        <Header />
        <Example />
      </main>
    </div>
  );
}

export async function getServerSideProps({ req }) {
  const cookies = parseCookies(req);

  return {
    props: {
      cookies: cookies,
    },
  };
}
