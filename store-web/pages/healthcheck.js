
import fetch from 'isomorphic-unfetch'
import useSWR from 'swr'

const API_URL = '/api/v1/health'
async function fetcher() {
    const res = await fetch(API_URL)
    const json = await res.json()
    return json
}

function HealthCheck(){
    const { data, error } = useSWR('/repos/zeit/next.js', fetcher)
    if (error) return <div>failed to load</div>
    if (!data) return <div>loading...</div>
    return <div>Next stars: {data.message.Name}</div>
}

export default HealthCheck