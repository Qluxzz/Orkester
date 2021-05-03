import { useState } from "react"
import { useHistory } from "react-router-dom"
import styled from "styled-components"


const Bar = styled.div`
    display: flex;
    flex-direction: column;
    padding: 10px;
`

const Input = styled.input`
    flex-grow: 1;
`

export default function SearchBar() {
    const [query, setQuery] = useState<string>()
    const history = useHistory()

    function updateQueryAndHistory(query: string) {
        setQuery(query)
        history.replace(`/search/${query}`)
    }

    return <Bar>
        <Input
            onChange={e => updateQueryAndHistory(e.target.value)}
            value={query}
        />
    </Bar>
}

