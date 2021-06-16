import { useState } from "react"
import { useHistory, useParams } from "react-router-dom"
import styled from "styled-components"


const Bar = styled.div`
    display: flex;
    flex-direction: column;
`

const Input = styled.input`
    flex-grow: 1;
`

export default function SearchBar() {
    const { query: initalQuery } = useParams<{ query: string }>()
    const [query, setQuery] = useState<string>(initalQuery)
    const history = useHistory()

    function updateQueryAndHistory(query: string) {
        setQuery(query)
        showSearchResults()
    }

    function showSearchResults() {
        history.replace(`/search/${query}`)
    }

    return <Bar>
        <Input
            onChange={e => updateQueryAndHistory(e.target.value)}
            value={query}
            onFocus={showSearchResults}
        />
    </Bar>
}

