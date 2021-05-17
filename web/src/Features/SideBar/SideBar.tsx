import React from "react"
import { Link } from "react-router-dom"
import styled from "styled-components"

const Container = styled.aside`
    --padding: 10px;

    background: #333;
    width: calc(150px - var(--padding));
    padding: var(--padding);
`


export default function SideBar() {
    return <Container>
        <Link to="/playlist/liked">Liked tracks</Link>
    </Container>
}