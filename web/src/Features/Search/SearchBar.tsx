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
    const { query } = useParams<{ query: string }>()
    const history = useHistory()

    function search(query: string) {
        history.replace(`/search/${query}`)
    }

    const searchQuery = query === undefined
        ? ""
        : query

    return <Bar>
        <Input
            onChange={e => search(e.target.value)}
            value={searchQuery}
            onFocus={() => {
                if (searchQuery.length > 0)
                    search(searchQuery)
            }}
        />
    </Bar>
}

