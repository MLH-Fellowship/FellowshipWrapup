import { Container } from "../../styles/fellow.style";
import Navbar from "../../components/layout/Navbar";
import Header from "../../components/containers/Fellow/Header";
import ProjectDetails from "../../components/containers/Fellow/ProjectDetails";

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

      <Container>
        <Header />
        <ProjectDetails />
      </Container>
    </>
  );
}
