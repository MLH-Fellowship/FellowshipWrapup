import Head from "next/head";
import fetch from "isomorphic-fetch";

import Header from "../../components/containers/Fellow/Header";
import ProjectDetails from "../../components/containers/Fellow/ProjectDetails";
import Milestones from "../../components/containers/Fellow/Milestones";
import Map from "../../components/containers/Fellow/Map";
import ProgressLayout from "../../components/containers/Fellow/ProgressTracker/ProgressLayout";
import Footer from "../../components/containers/Fellow/Footer/Footer";

const Fellow = ({ accountInfo, issueInfo, contributedTo }) => {
  // console.log(accountInfo.User);
  // console.log(issueInfo.User.Issues.Nodes);

  const filteredIssues = issueInfo.User.Issues.Nodes.filter((el) =>
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
          contributions={contributedTo.User.PullRequests.Nodes}
        />
        <Milestones issues={filteredIssues} />
        <ProgressLayout />
        <Footer />
      </div>
    </>
  );
};

Fellow.getInitialProps = async ({ query }) => {
  const [resAcc, resIss, contributedTo] = await Promise.all([
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
  ]);

  return {
    accountInfo: resAcc,
    issueInfo: resIss,
    contributedTo,
    query,
  };
};

export default Fellow;
