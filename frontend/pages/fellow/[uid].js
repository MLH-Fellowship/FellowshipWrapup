import Head from "next/head";
import Router from "next/router";
import fetch from "isomorphic-fetch";

import Header from "../../components/containers/Fellow/Header";
import ProjectDetails from "../../components/containers/Fellow/ProjectDetails";
import Milestones from "../../components/containers/Fellow/Milestones";
import Map from "../../components/containers/Fellow/Map";
import ProgressLayout from "../../components/containers/Fellow/ProgressTracker/ProgressLayout";
import FellowMap from "../../components/containers/Fellow/FellowMap/index";
import Footer from "../../components/containers/Fellow/Footer/Footer";

const Fellow = ({
  accountInfo,
  issueInfo,
  contributedTo,
  commits,
  prContributions,
  prs,
}) => {
  if (!accountInfo) {
    return <h1>The account you are looking for doesn't seem to exist</h1>;
  }
  const filteredIssues = issueInfo.User.Issues.Nodes.filter((el) =>
    el.Url.startsWith("https://github.com/MLH-Fellowship/")
  );

  const filteredContributions = contributedTo.User.PullRequests.Nodes.filter(
    (el) => el.Url.startsWith("https://github.com/MLH-Fellowship/")
  );

  const filteredCommits = commits.User.PullRequests.Nodes.filter((el) =>
    el.Url.startsWith("https://github.com/MLH-Fellowship/")
  );

  const filteredPrContributions = prContributions.User.PullRequests.Nodes.filter(
    (el) => el.Url.startsWith("https://github.com/MLH-Fellowship/")
  );

  const filteredPrs = prs.User.PullRequests.Nodes.filter((el) =>
    el.Url.startsWith("https://github.com/MLH-Fellowship/")
  );

  return (
    <>
      <Head>
        <title>{accountInfo.User.Name} | MLH Fellow</title>
      </Head>
      <nav className="navbar navbar-expand-sm navbar-light">
        <ul className="navbar-nav">
          <li className="nav-item">
            <a className="nav-link" href="#">
              {accountInfo.User.Name}
            </a>
          </li>
        </ul>
      </nav>
      <div className="container">
        <Header accountInfo={accountInfo} />
        <Map />
        <ProjectDetails
          accountInfo={accountInfo.User}
          contributions={filteredContributions}
        />
        <Milestones
          issues={filteredIssues}
          commits={filteredCommits}
          prContributions={filteredPrContributions}
        />
        <ProgressLayout issues={filteredIssues} prs={filteredPrs} />
        <FellowMap />
        <Footer />
      </div>
    </>
  );
};

Fellow.getInitialProps = async ({ query }) => {
  const [
    resAcc,
    resIss,
    contributedTo,
    commits,
    prContributions,
    prs,
  ] = await Promise.all([
    fetch(`${process.env.BACKEND_URL}/accountinfo/${query.uid}`, {
      method: "POST",
      body: JSON.stringify({
        secret: `${process.env.BACKEND_SECRET}`,
      }),
    }).then((res) => res.json()),
    fetch(`${process.env.BACKEND_URL}/issuescreated/${query.uid}`, {
      method: "POST",
      body: JSON.stringify({
        secret: `${process.env.BACKEND_SECRET}`,
      }),
    }).then((res) => res.json()),
    fetch(`${process.env.BACKEND_URL}/repocontributedto/${query.uid}`, {
      method: "POST",
      body: JSON.stringify({
        secret: `${process.env.BACKEND_SECRET}`,
      }),
    }).then((res) => res.json()),
    fetch(`${process.env.BACKEND_URL}/pullrequestcommits/${query.uid}`, {
      method: "POST",
      body: JSON.stringify({
        secret: `${process.env.BACKEND_SECRET}`,
      }),
    }).then((res) => res.json()),
    fetch(`${process.env.BACKEND_URL}/prcontributions/${query.uid}`, {
      method: "POST",
      body: JSON.stringify({
        secret: `${process.env.BACKEND_SECRET}`,
      }),
    }).then((res) => res.json()),
    fetch(`${process.env.BACKEND_URL}/pullrequests/${query.uid}`, {
      method: "POST",
      body: JSON.stringify({
        secret: `${process.env.BACKEND_SECRET}`,
      }),
    }).then((res) => res.json()),
  ]);

  return {
    accountInfo: resAcc,
    issueInfo: resIss,
    contributedTo,
    commits,
    prContributions,
    prs,
  };
};

export default Fellow;
