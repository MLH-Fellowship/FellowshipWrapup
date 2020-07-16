import Header from "../../components/containers/Fellow/Header";
import ProjectDetails from "../../components/containers/Fellow/ProjectDetails";
import Milestones from "../../components/containers/Fellow/Milestones";
import Map from "../../components/containers/Fellow/Map";

export default function Fellow() {
  return (
    <>
      <nav class="navbar navbar-expand-sm navbar-light">
        <ul class="navbar-nav">
          <li class="nav-item">
            <a class="nav-link" href="#">
              Sebastian Crossa
            </a>
          </li>
        </ul>
      </nav>
      <div className="container">
        <Header />
        <Map />
        <ProjectDetails />
        <Milestones />
      </div>
    </>
  );
}
