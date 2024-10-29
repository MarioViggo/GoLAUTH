import React, {useState} from "react";
import { Col, Row } from "react-bootstrap";
import { Navigate } from 'react-router';
import { Container } from "./styles";

import Axios from '../../helpers/Axios';
import Input from "../../components/input";

export default function SignIn() {
    const [email, setEmail] = useState(null);
    const [password, setPassword] = useState(null);

    const token = localStorage.getItem('token')

    const signIn = async () => {
        await Axios.post("/auth/signin", { email: email, password: password }).then(res => {
            localStorage.setItem('token', res.data.Token);
            window.location.reload();
        }).catch(err => alert("Credenciais incorretas!"))
    }

    if (!token) {
        return (
            <Container className="container-lg">
                <Row>
                    <Col md={15}>
                        <Input
                            name="email"
                            label="Email"
                            type="email"
                            value={email}
                            placeholder="Digite seu Email"
                            onChange={(e) => setEmail(e.target.value)}
                        />
                    </Col>
                </Row>
                <Row>
                    <Col md={15}>
                        <Input
                            name="password"
                            label="Senha"
                            type="password"
                            value={password}
                            placeholder="Digite sua senha"
                            onChange={(e) => setPassword(e.target.value)}
                        />
                    </Col>
                </Row>
                <Row>
                    <Col md={4}>
                        <Row className="mt-4">
                            <Col md={{ span: 2, offset: 4 }}>
                                <button onClick={signIn} class="btn btn-primary">
                                    Register
                                </button>
                            </Col>

                        </Row>
                    </Col>
                </Row>
            </Container>

        )
    } else {
        alert("Você está logado!")
        return <Navigate to="/" />
    }
}