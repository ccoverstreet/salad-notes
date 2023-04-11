import { c as create_ssr_component, d as add_attribute, v as validate_component } from "../../chunks/index.js";
const MDViewer = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let { content } = $$props;
  let display = {};
  if ($$props.content === void 0 && $$bindings.content && content !== void 0)
    $$bindings.content(content);
  display.innerHTML = content;
  return `<div id="${"output"}"${add_attribute("this", display, 0)}></div>`;
});
const MDEditor_svelte_svelte_type_style_lang = "";
const css = {
  code: "#wrapper.svelte-1wcsfwt.svelte-1wcsfwt{height:60vh;width:100%;display:flex;grid-template-columns:50% 50%;flex-wrap:wrap;justify-content:center}#wrapper.svelte-1wcsfwt>.svelte-1wcsfwt{flex-grow:1;min-width:30rem}#wrapper.svelte-1wcsfwt>textarea.svelte-1wcsfwt{tab-size:4}",
  map: null
};
const MDEditor = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  let inputDiv;
  let renderedInput;
  $$result.css.add(css);
  return `<div id="${"wrapper"}" class="${"svelte-1wcsfwt"}"><textarea class="${"svelte-1wcsfwt"}"${add_attribute("this", inputDiv, 0)}>${""}</textarea>
	${validate_component(MDViewer, "MDViewer").$$render(
    $$result,
    {
      id: "viewer",
      content: renderedInput,
      style: ""
    },
    {},
    {}
  )}
</div>`;
});
const Page = create_ssr_component(($$result, $$props, $$bindings, slots) => {
  return `<div style="${"display:flex; justify-content: center;"}"><h2>SaladNotes</h2></div>
<div>${validate_component(MDEditor, "MDEditor").$$render($$result, {}, {}, {})}</div>`;
});
export {
  Page as default
};
