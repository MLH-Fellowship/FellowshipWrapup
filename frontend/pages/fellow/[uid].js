import Head from "next/head";
import { getUserInfo } from "../../utils/user-info";

import Header from "../../components/containers/Fellow/Header";
import ProjectDetails from "../../components/containers/Fellow/ProjectDetails";
import Milestones from "../../components/containers/Fellow/Milestones";
import Map from "../../components/containers/Fellow/Map";
import ProgressLayout from "../../components/containers/Fellow/ProgressTracker/ProgressLayout";
import Footer from "../../components/containers/Fellow/Footer/Footer";

const Fellow = ({ accountInfo }) => {
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
        <ProjectDetails />
        <Milestones />
        <ProgressLayout />
        <Footer />
      </div>
    </>
  );
};

Fellow.getInitialProps = async () => {
  const info = getUserInfo();

  return {
    accountInfo: info,
  };
};

export default Fellow;
