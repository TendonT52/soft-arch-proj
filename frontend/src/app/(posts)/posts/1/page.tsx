import { type InitialEditorStateType } from "@lexical/react/LexicalComposer";
import { PostEditor } from "@/components/post-editor";

/* DUMMY */
type Post = {
  topic: string;
  description: InitialEditorStateType;
};

// prettier-ignore
const post: Post = {
  topic: "Social Engineering",
  description: JSON.stringify({"root":{"children":[{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"What to expect from here on out","type":"text","version":1}],"direction":"ltr","format":"start","indent":0,"type":"heading","version":1,"tag":"h2"},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"What follows from here is just a bunch of absolute ","type":"text","version":1},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"nonsense","type":"text","version":1}],"direction":"ltr","format":"","indent":0,"type":"link","version":1,"rel":"noreferrer","target":null,"title":null,"url":"https://tailwindcss.com"},{"detail":0,"format":0,"mode":"normal","style":"","text":" I've written to dogfood the plugin itself. It includes every sensible typographic element I could think of, like ","type":"text","version":1},{"detail":0,"format":1,"mode":"normal","style":"","text":"bold text","type":"text","version":1},{"detail":0,"format":0,"mode":"normal","style":"","text":", unordered lists, ordered lists, code blocks, block quotes, ","type":"text","version":1},{"detail":0,"format":2,"mode":"normal","style":"","text":"and even italics","type":"text","version":1},{"detail":0,"format":0,"mode":"normal","style":"","text":".","type":"text","version":1}],"direction":"ltr","format":"start","indent":0,"type":"paragraph","version":1},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"It's important to cover all of these use cases for a few reasons:","type":"text","version":1}],"direction":"ltr","format":"start","indent":0,"type":"paragraph","version":1},{"children":[{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"We want everything to look good out of the box.","type":"text","version":1}],"direction":"ltr","format":"","indent":0,"type":"listitem","version":1,"value":1},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"Really just the first reason, that's the whole point of the plugin.","type":"text","version":1}],"direction":"ltr","format":"","indent":0,"type":"listitem","version":1,"value":2},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"Here's a third pretend reason though a list with three items looks more realistic than a list with two items.","type":"text","version":1}],"direction":"ltr","format":"","indent":0,"type":"listitem","version":1,"value":3}],"direction":"ltr","format":"","indent":0,"type":"list","version":1,"listType":"number","start":1,"tag":"ol"},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"Now we're going to try out another header style.","type":"text","version":1}],"direction":"ltr","format":"start","indent":0,"type":"paragraph","version":1},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"Typography should be easy","type":"text","version":1}],"direction":"ltr","format":"start","indent":0,"type":"heading","version":1,"tag":"h3"},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"So that's a header for you — with any luck if we've done our job correctly that will look pretty reasonable.","type":"text","version":1}],"direction":"ltr","format":"start","indent":0,"type":"paragraph","version":1},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"Something a wise person once told me about typography is:","type":"text","version":1}],"direction":"ltr","format":"start","indent":0,"type":"paragraph","version":1},{"children":[{"detail":0,"format":0,"mode":"normal","style":"","text":"Typography is pretty important if you don't want your stuff to look like trash. Make it good then it won't be bad.","type":"text","version":1}],"direction":"ltr","format":"","indent":0,"type":"quote","version":1}],"direction":"ltr","format":"","indent":0,"type":"root","version":1}}),
};
/* DUMMY */

export default function Page() {
  return (
    <PostEditor
      post={post}
      editable={false}
    />
  );
}
