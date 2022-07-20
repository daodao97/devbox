import { ElIcon } from 'element-plus'
import { isFunction, isString } from 'lodash'
import { Component } from 'vue'
export * from './SvgIcon'

export const Render = (props: { content : Component | Function | String | unknown}) => {
  if (isFunction(props.content)) {
    return (props.content as Function)()
  }
  if (isString(props.content)) {
    return <span>{ props.content }</span>
  }
  const tmp = props.content
  return <tmp/>
}

export const IconWrap = (icon: Component ) => {
  const tmp = icon
  return <ElIcon><tmp/></ElIcon>
}

