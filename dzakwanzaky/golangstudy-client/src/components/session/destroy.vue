<template lang='pug'>
    h3 Logging Out!
</template>

<script>

import Cookies from 'js-cookie'

export default {
    name : 'session-destroy',
    beforeRouteEnter(to,from,next){
        Cookies.remove('golangstudy')

        location.reload(true)
    }
}
</script>
