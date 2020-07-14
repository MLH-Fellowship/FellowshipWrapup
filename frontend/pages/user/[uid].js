import { useRouter } from "next/router";

import Header from "../../components/User/Header";

export default function User() {
  const router = useRouter();

  return (
    <div>
      <Header />
    </div>
  );
}
