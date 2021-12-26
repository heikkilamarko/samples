export function processPerson(item) {
  return {
    id: Number(item.id),
    name: item.name,
    age: Number(item.age),
    height: Number(item.height),
    is_active: item.is_active === "true",
    created_at: new Date(item.created_at),
  };
}
